package minequery

// defaultMinecraftPort is a default port Minecraft server runs on and which
// will be used when server port is left as zero value.
const defaultMinecraftPort = 25565

// pingFunction is a generic function type for version-specific ping functions to conform to.
type pingFunction func(string, int) (interface{}, error)

// pingGeneric accepts version-specific ping function and host/port pair. Then it performs
// (if necessary, see PreferSRVRecord) SRV lookup, and attempts to use the SRV record hostname and port
// (first returned record is used if more than one is returned, see net.LookupSRV documentation)
// to ping, if lookup fails or ping fails, the provided hostname/port pair is used directly.
func (p *Pinger) pingGeneric(
	pingFn pingFunction,
	host string, port int,
) (interface{}, error) {
	// Use default Minecraft port if port is 0
	if port == 0 {
		port = defaultMinecraftPort
	}

	if p.PreferSRVRecord {
		// When SRV record is preferred, try resolving it
		srvHost, srvPort, err := resolveSRV(host)
		if err != nil {
			// If not UseStrict, continue pinging on the desired host/port
			if p.UseStrict {
				// If UseStrict, SRV lookup error is fatal
				return nil, err
			}

		} else {
			// If SRV lookup is successful, check if there are any records
			if srvHost != "" {
				status, err := pingFn(srvHost, int(srvPort))
				if err == nil {
					// Success, SRV record ping passed
					return status, nil
				}

				// If pinging on the SRV record failed and UseStrict is set,
				// this is fatal enough to raise an error
				if p.UseStrict {
					return nil, err
				}
			}
		}
	}

	// Otherwise just ping normally
	return pingFn(host, port)
}

// Defines a message structure for encapsulating metadata and payload
package communication

// Message encapsulates metadata and payload.
type Message struct {
	Data      []byte // Encoded/Modulated data
	Protocols struct {
		Codification string
		Modulation   string
	}
}

package cookiejar

import (
	"encoding/json"
	"golang.org/x/net/publicsuffix"
)

// TODO: cookiejar.Jar vs Browser. 1. example.com: Set-Cookie: id=123; Path=/ -> 2. example.com: Set-Cookie: id=456; Domain=example.com; Path=/

func NewWithPublicSuffix() (*Jar, error) {
	return New(&Options{
		PublicSuffixList: publicsuffix.List,
	})
}

func NewFromJSON(jsonStr string) (*Jar, error) {
	jar, err := NewWithPublicSuffix()
	if err != nil {
		return nil, err
	}
	return jar, jar.UnmarshalJSON([]byte(jsonStr))
}

func (j *Jar) UnmarshalJSON(data []byte) error {
	j.mu.Lock()
	defer j.mu.Unlock()

	var importData struct {
		NextSeqNum uint64                      `json:"nextSeqNum"`
		Entries    map[string]map[string]entry `json:"entries"`

		ImportSessionCookies bool `json:"importSessionCookies,omitempty"`
	}

	if err := json.Unmarshal(data, &importData); err != nil {
		return err
	}

	if importData.ImportSessionCookies {
		j.entries = importData.Entries
	} else {
		filteredEntries := make(map[string]map[string]entry)
		for host, submap := range importData.Entries {
			filteredSubmap := make(map[string]entry)
			for key, value := range submap {
				if value.Persistent {
					filteredSubmap[key] = value
				}
			}
			if len(filteredSubmap) > 0 {
				filteredEntries[host] = filteredSubmap
			}
		}
		j.entries = filteredEntries
	}

	j.nextSeqNum = importData.NextSeqNum

	return nil
}

func (j *Jar) MarshalJSON() ([]byte, error) {
	j.mu.Lock()
	defer j.mu.Unlock()

	exportData := map[string]interface{}{
		"nextSeqNum": j.nextSeqNum,
		"entries":    j.entries,
	}

	return json.Marshal(exportData)
}

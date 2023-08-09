package usenet

import "github.com/chrisfarms/nntp"

func FindGroup(c *nntp.Conn, groups []string) error {
	var err error
	for _, g := range groups {
		_, _, _, err = c.Group(g)
		if err == nil {
			return nil
		}
	}
	return err
}

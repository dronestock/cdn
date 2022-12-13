package main

func (p *plugin) refresh() (undo bool, err error) {
	for _, domain := range p.domains {
		switch domain.Provider {
		case providerTencentyun:
			err = p.tencentyun(domain)
		}
	}

	return
}

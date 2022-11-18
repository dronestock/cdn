package main

func (p *plugin) refresh() (undo bool, err error) {
	switch p.Provider {
	case providerTencentyun:
		err = p.tencentyun()
	}

	return
}

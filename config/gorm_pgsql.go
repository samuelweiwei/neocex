package config

type Pgsql struct {
	GeneralDB
}

func (p *Pgsql) Dsn(dbname string) string {
	return "host=" + p.Path + " user=" + p.Username + " password=" + p.Password + " dbname=" + dbname + " port=" + p.Port + " " + p.Config
}

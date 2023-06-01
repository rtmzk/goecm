package options

func (o *Options) Validate() []error {
	var errs []error

	errs = append(errs, o.GenericServerRunOptions.Validate()...)
	errs = append(errs, o.InsecureServing.Validate()...)
	errs = append(errs, o.SecureServing.Validate()...)
	//errs = append(errs, o.MySQLOptions.Validate()...)
	errs = append(errs, o.SQLiteOptions.Validate()...)
	//errs = append(errs, o.JwtOptions.Validate()...)
	errs = append(errs, o.Log.Validate()...)

	return errs
}

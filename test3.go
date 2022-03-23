
func (r schemaResolver) ViewerSettings(ctx context.Context) (*settingsCascade, error) {
	user, err := CurrentUser(ctx, r.db)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return &settingsCascade{db: r.db, unauthenticatedActor: true}, nil
	}
	return &settingsCascade{db: r.db, subject: &settingsSubject{user: user}}, nil
}

// Deprecated: in the GraphQL API
func (r *schemaResolver) ViewerConfiguration(ctx context.Context) (*settingsCascade, error) {
	return schemaResolver{db: r.db}.ViewerSettings(ctx)
}

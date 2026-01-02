package seeds

import "github.com/lwmacct/260101-go-pkg-ddd/ddd/core/infrastructure/database"

// DefaultSeeders returns the default ordered seeders that bootstrap the system.
// Keep RBAC first because it provisions permissions/roles required by other seeders.
// SettingCategorySeeder must run before SettingSeeder to ensure categories exist.
// OrganizationSeeder runs last because it depends on UserSeeder (admin user).
func DefaultSeeders() []database.Seeder {
	return []database.Seeder{
		&RBACSeeder{},
		&UserSeeder{},
		&SettingCategorySeeder{},
		&SettingSeeder{},
		&OrganizationSeeder{},
	}
}

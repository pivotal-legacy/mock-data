package main

import (
	"github.com/spf13/viper"
	"testing"
)

// create fake database for testing database.go
func createFakeTableForDatabaseTest() {
	setDatabaseConfigForTest()
	postgresOrGreenplum()
	ExecuteDemoDatabase()
	_, err := ExecuteDB("CREATE TABLE testme (id int);") // AT least create one table
	if err != nil {
		Fatalf("TestExecuteDemoDatabasePreCleanup, failed to create a demo table, err: %v", err)
	}
}

// Test: executeDemoDatabasePreCleanup, checking for script cleanup ability
func TestExecuteDemoDatabasePreCleanup(t *testing.T) {
	createFakeTableForDatabaseTest()
	tests := []struct {
		name string
		f    bool
		want int
	}{
		{"environment_variable_off", false, 1},
		{"environment_variable_on", true, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			viper.Set("MOCK_DATA_TEST_RUNNER", tt.f)
			executeDemoDatabasePreCleanup()
			got := allTablesPostgres("")
			if len(got) < tt.want {
				t.Errorf("TestExecuteDemoDatabasePreCleanup = %v, want %v", len(got), tt.want)
			}
		})
	}
	viper.Set("MOCK_DATA_TEST_RUNNER", true) // set this for the reminder of the test
}

// Test: ExecuteDemoDatabase, check if all the tables are created
func TestExecuteDemoDatabase(t *testing.T) {
	createFakeTableForDatabaseTest()
	t.Run("should_extract_all_tables", func(t *testing.T) {
		if got := allTablesPostgres(""); len(got) < 15 {
			t.Errorf("TestExecuteDemoDatabase = %v, want >= %d tables", len(got), 15)
		}
	})
}

// Test: MockDatabase, Mock the entire database
func TestMockDatabase(t *testing.T) {
	createFakeTableForDatabaseTest()
	cmdOptions.Rows = 100
	MockDatabase()
	tests := []struct {
		name  string
		table string
	}{
		{"mock_database_check_actor_table", "public.actor"},
		{"mock_database_check_address_table", "public.address"},
		{"mock_database_check_category_table", "public.category"},
		{"mock_database_check_city_table", "public.city"},
		{"mock_database_check_country_table", "public.country"},
		{"mock_database_check_customer_table", "public.customer"},
		{"mock_database_check_film_table", "public.film"},
		{"mock_database_check_film_actor_table", "public.film_actor"},
		{"mock_database_check_film_category_table", "public.film_category"},
		{"mock_database_check_inventory_table", "public.inventory"},
		{"mock_database_check_language_table", "public.language"},
		{"mock_database_check_payment_table", "public.payment"},
		{"mock_database_check_rental_table", "public.rental"},
		{"mock_database_check_staff_table", "public.staff"},
		{"mock_database_check_store_table", "public.store"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Why > 0, since we do delete some rows with deleteViolatingConstraintKeys, so we dont know
			// how much we deleted, so lets just say it should have greater than 0 rows
			if got := TotalRows(tt.table); got <= 0 {
				t.Errorf("TestMockDatabase = %v, want %v", got, "> 0")
			}
		})
	}
}

// Test: dbExtractTables, check if all the tables are returned
func TestDbExtractTables(t *testing.T) {
	createFakeTableForDatabaseTest()
	t.Run("should_return_all_tables", func(t *testing.T) {
		if got := dbExtractTables(""); len(got) < 15 {
			t.Errorf("TestDbExtractTables = %v, want >= %d tables", len(got), 15)
		}
	})
}

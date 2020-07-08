package model

import (
	"database/sql"
	"time"
)

const createTables = `
CREATE TABLE IF NOT EXISTS "permission" (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS "role" (
    id SERIAL PRIMARY KEY,
    name VARCHAR(30) UNIQUE NOT NULL,
    description TEXT
);

CREATE TABLE IF NOT EXISTS "permission_role" (
    id_permission INTEGER REFERENCES permission(id) ON DELETE CASCADE,
    id_role INTEGER REFERENCES role(id) ON DELETE CASCADE,
    PRIMARY KEY (id_permission, id_role)
);

CREATE TABLE IF NOT EXISTS "user_auth" (
   id SERIAL PRIMARY KEY,
   user_name VARCHAR(30) UNIQUE NOT NULL,
   password TEXT NOT NULL,
   id_role INTEGER REFERENCES role(id) ON DELETE SET NULL
);

CREATE TABLE IF NOT EXISTS "user" (
   id SERIAL PRIMARY KEY,
   name VARCHAR(30) NOT NULL,
   last_name VARCHAR(50) NOT NULL,
   identification VARCHAR(20) NOT NULL UNIQUE,
   birthday DATE
);

CREATE TABLE IF NOT EXISTS "client" (
    user_id INTEGER REFERENCES "user"(id) ON DELETE CASCADE,
    address VARCHAR(80) NOT NULL,
    PRIMARY KEY (user_id)
);

CREATE TABLE IF NOT EXISTS "employee" (
    user_id INTEGER REFERENCES "user"(id) ON DELETE CASCADE,
    salary FLOAT NOT NULL,
    PRIMARY KEY (user_id)
);

CREATE TABLE IF NOT EXISTS "service" (
    id SERIAL PRIMARY KEY,
    name VARCHAR(30) UNIQUE NOT NULL,
    description text NOT NULL,
    price FLOAT NOT NULL,
    hour INT NOT NULL,
    minute INT NOT NULL
);

CREATE TABLE IF NOT EXISTS "employee_service" (
    id_service INTEGER REFERENCES service(id) ON DELETE CASCADE,
    id_employee INTEGER REFERENCES employee(user_id) ON DELETE CASCADE,
    PRIMARY KEY (id_employee, id_service)
);

CREATE TABLE IF NOT EXISTS "branch_office" (
    id SERIAL PRIMARY KEY,
    city VARCHAR(50) NOT NULL,
    province VARCHAR(50) NOT NULL,
    address VARCHAR(50) NOT NULL,
    check_in_time TIME NOT NULL,
    exit_time TIME NOT NULL
);

CREATE TABLE IF NOT EXISTS "branch_office_service" (
    id_service INTEGER REFERENCES service(id) ON DELETE CASCADE,
    id_branch_office INTEGER REFERENCES branch_office(id) ON DELETE CASCADE,
    PRIMARY KEY (id_branch_office, id_service)
);

CREATE TABLE IF NOT EXISTS "appointment" (
    id SERIAL PRIMARY KEY,
    total_price FLOAT NOT NULL,
    start_time TIMESTAMP NOT NULL,
    end_time TIMESTAMP NOT NULL,
    id_client INTEGER REFERENCES client(user_id) ON DELETE SET NULL,
    id_branch_office INTEGER REFERENCES branch_office(id) ON DELETE SET NULL
);

CREATE TABLE IF NOT EXISTS "employee_appointment" (
    id_employee INTEGER REFERENCES employee(user_id) ON DELETE CASCADE,
    id_appointment INTEGER REFERENCES appointment(id) ON DELETE CASCADE,
    PRIMARY KEY (id_employee, id_appointment)
)`

type Permission struct {
	Id uint64 `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}
type Role struct {
	Id uint64 `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
	Description sql.NullString `db:"description" json:"description"`
}
type PermissionRole struct {
	Id uint64 `db:"id"`
	IdPermission uint64 `db:"id_permission"`
	IdRole uint64 `db:"id_role"`
}
type UserAuth struct {
	Id uint64 `db:"id"`
	Username string `db:"user_name"`
	Password string `db:"password"`
	IdRole uint64 `db:"id_role"`
}
type User struct {
	Id uint64 `db:"id"`
	Name string `db:"name"`
	LastName string `db:"last_name"`
	Identification string `db:"identification"`
	Birthday sql.NullTime `db:"birthday"`
}
type Client struct {
	Id uint64 `db:"user_id"`
	Address string `db:"address"`
}
type Employee struct {
	Id uint64 `db:"user_id"`
	Salary float64 `db:"salary"`
}
type Service struct {
	Id uint64 `db:"id"`
	Name string `db:"name"`
	Description string `db:"description"`
	Price float64 `db:"price"`
	Hour uint8 `db:"hour"`
	Minute uint8 `db:"minute"`
}
type EmployeeService struct {
	Id uint64 `db:"id"`
	IdEmployee uint64 `db:"id_employee"`
	IdService uint64 `db:"id_service"`
}
type BranchOffice struct {
	Id uint64 `db:"id" json:"id"`
	City string `db:"city" json:"city"`
	Province string `db:"province" json:"province"`
	Address string `db:"address" json:"address"`
	CheckInTime string `db:"check_in_time" json:"checkInTime"`
	ExitTime string `db:"exit_time" json:"exitTime"`
}
type BranchOfficeService struct {
	Id uint64 `db:"id"`
	IdBranchOffice uint64 `db:"id_branch_office"`
	IdService uint64 `db:"id_service"`
}
type Appointment struct {
	Id uint64 `db:"id"`
	TotalPrice float64 `db:"total_price"`
	StartTime time.Time `db:"start_time"`
	EndTime time.Time `db:"end_time"`
	IdClient uint64 `db:"id_client"`
	IdBranchOffice uint64 `db:"id_branch_office"`
}

type EmployeeAppointment struct {
	Id uint64 `db:"id"`
	IdAppointment uint64 `db:"id_appointment"`
	IdEmployee uint64 `db:"id_employee"`
}
func Migrate() error {
	db, err := ConnectionDatabase()
	if err != nil { return err }
	db.MustExec(createTables)
	defer db.Close()
	return nil
}
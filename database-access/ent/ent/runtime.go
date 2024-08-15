// Code generated by ent, DO NOT EDIT.

package ent

import (
	"demo/ent/course"
	"demo/ent/department"
	"demo/ent/enrollment"
	"demo/ent/instructor"
	"demo/ent/student"
	"demo/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	courseFields := schema.Course{}.Fields()
	_ = courseFields
	// courseDescTitle is the schema descriptor for title field.
	courseDescTitle := courseFields[0].Descriptor()
	// course.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	course.TitleValidator = courseDescTitle.Validators[0].(func(string) error)
	departmentFields := schema.Department{}.Fields()
	_ = departmentFields
	// departmentDescName is the schema descriptor for name field.
	departmentDescName := departmentFields[0].Descriptor()
	// department.NameValidator is a validator for the "name" field. It is called by the builders before save.
	department.NameValidator = departmentDescName.Validators[0].(func(string) error)
	enrollmentFields := schema.Enrollment{}.Fields()
	_ = enrollmentFields
	// enrollmentDescSemester is the schema descriptor for semester field.
	enrollmentDescSemester := enrollmentFields[0].Descriptor()
	// enrollment.SemesterValidator is a validator for the "semester" field. It is called by the builders before save.
	enrollment.SemesterValidator = enrollmentDescSemester.Validators[0].(func(string) error)
	// enrollmentDescYear is the schema descriptor for year field.
	enrollmentDescYear := enrollmentFields[1].Descriptor()
	// enrollment.YearValidator is a validator for the "year" field. It is called by the builders before save.
	enrollment.YearValidator = enrollmentDescYear.Validators[0].(func(int) error)
	instructorFields := schema.Instructor{}.Fields()
	_ = instructorFields
	// instructorDescName is the schema descriptor for name field.
	instructorDescName := instructorFields[0].Descriptor()
	// instructor.NameValidator is a validator for the "name" field. It is called by the builders before save.
	instructor.NameValidator = instructorDescName.Validators[0].(func(string) error)
	studentFields := schema.Student{}.Fields()
	_ = studentFields
	// studentDescName is the schema descriptor for name field.
	studentDescName := studentFields[0].Descriptor()
	// student.NameValidator is a validator for the "name" field. It is called by the builders before save.
	student.NameValidator = studentDescName.Validators[0].(func(string) error)
}
package students

import (
	"github.com/SocBongDev/soc-bong/internal/common"
	"github.com/SocBongDev/soc-bong/internal/entities"
	"github.com/pocketbase/dbx"
)

func mapStudents(dbResult []DbStudentResult) []Student {
	studentsMap := make(map[int]*Student)
	for _, v := range dbResult {
		if _, ok := studentsMap[v.Id]; ok {
			studentsMap[v.Id] = &Student{
				BaseEntity: common.BaseEntity{
					Id:        v.Id,
					CreatedAt: v.CreatedAt,
					UpdatedAt: v.UpdatedAt,
				},
				Class: entities.Class{
					BaseEntity: common.BaseEntity{
						Id:        v.ClassId,
						CreatedAt: v.ClassCreatedAt,
						UpdatedAt: v.ClassUpdatedAt,
					},
					WriteClassRequest: entities.WriteClassRequest{
						Name:      v.ClassName,
						Grade:     v.ClassGrade,
						AgencyId:  v.AgencyId,
						TeacherId: v.TeacherId,
					},
				},
				WriteStudentRequest: WriteStudentRequest{
					BirthPlace:               v.BirthPlace,
					Dob:                      v.Dob,
					EnrolledAt:               v.EnrolledAt,
					Ethnic:                   v.Ethnic,
					FirstName:                v.FirstName,
					Gender:                   v.Gender,
					LastName:                 v.LastName,
					PermanentAddressCommune:  v.PermanentAddressCommune,
					PermanentAddressDistrict: v.PermanentAddressDistrict,
					PermanentAddressProvince: v.PermanentAddressProvince,
					TempAddress:              v.TempAddress,
					AgencyId:                 v.AgencyId,
					ClassId:                  v.ClassId,
					FatherBirthPlace:         v.FatherBirthPlace,
					MotherBirthPlace:         v.MotherBirthPlace,
					FatherDob:                v.FatherDob,
					MotherDob:                v.MotherDob,
					FatherName:               v.FatherName,
					MotherName:               v.MotherName,
					Landlord:                 v.Landlord,
					FatherOccupation:         v.FatherOccupation,
					MotherOccupation:         v.MotherOccupation,
					FatherPhoneNumber:        v.FatherPhoneNumber,
					MotherPhoneNumber:        v.MotherPhoneNumber,
					ResRegistration:          v.ResRegistration,
					Roi:                      v.Roi,
					Zalo:                     v.Zalo,
				},
			}
		}
	}
	students := make([]Student, 0, len(studentsMap))
	for _, s := range studentsMap {
		students = append(students, *s)
	}

	return students
}

func allHook(q *dbx.Query, sliceA any, op func(sliceB any) error) error {
	switch v := sliceA.(type) {
	case *[]Student:
		dbStudentResults := new([]DbStudentResult)
		if err := op(dbStudentResults); err != nil {
			return err
		}

		students := mapStudents(*dbStudentResults)
		*v = students
		return nil
	default:
		return op(sliceA)
	}
}

package patients



import (
	"github.com/gin-gonic/gin"
)

type PatientSerializer struct {
	c *gin.Context
	PatientModel
}

type PatientsSerializer struct {
	C     *gin.Context
	Patients []PatientModel
}

type PatientResponse struct {
	PatientModel
}

func NewPatientSerializer(c *gin.Context, Patient PatientModel) PatientSerializer {
	return PatientSerializer{
		c:         c,
		PatientModel: Patient,
	}
}

func NewPatientsSerializer(c *gin.Context, Patients []PatientModel) PatientsSerializer {
	return PatientsSerializer{
		C:     c,
		Patients: Patients,
	}
}

func (s *PatientSerializer) Response() PatientResponse {
	return PatientResponse{
		PatientModel: s.PatientModel,
	}
}

func (s *PatientsSerializer) Response() []PatientResponse {
	response := make([]PatientResponse, len(s.Patients))
	for i, PatientModel := range s.Patients {
		serializer := NewPatientSerializer(s.C, PatientModel)
		response[i] = serializer.Response()
	}
	return response
}

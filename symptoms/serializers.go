package symptoms

import (
	"github.com/gin-gonic/gin"
)

type SymptomsSerializer struct {
	c *gin.Context
	SymptomsModel
}

type SymptomssSerializer struct {
	C         *gin.Context
	Symptomss []SymptomsModel
}

type SymptomsResponse struct {
	SymptomsModel
}

func NewSymptomsSerializer(c *gin.Context, Symptoms SymptomsModel) SymptomsSerializer {
	return SymptomsSerializer{
		c:             c,
		SymptomsModel: Symptoms,
	}
}

func NewSymptomssSerializer(c *gin.Context, Symptomss []SymptomsModel) SymptomssSerializer {
	return SymptomssSerializer{
		C:         c,
		Symptomss: Symptomss,
	}
}

func (s *SymptomsSerializer) Response() SymptomsResponse {
	return SymptomsResponse{
		SymptomsModel: s.SymptomsModel,
	}
}

func (s *SymptomssSerializer) Response() []SymptomsResponse {
	response := make([]SymptomsResponse, len(s.Symptomss))
	for i, SymptomsModel := range s.Symptomss {
		serializer := NewSymptomsSerializer(s.C, SymptomsModel)
		response[i] = serializer.Response()
	}
	return response
}

package illnesssymptom

import (
	"github.com/gin-gonic/gin"
)

type llnessSymptomsSerializer struct {
	c *gin.Context
	IllnessSymptomsModel
}

type llnessSymptomssSerializer struct {
	C               *gin.Context
	llnessSymptomss []IllnessSymptomsModel
}

type llnessSymptomsResponse struct {
	IllnessSymptomsModel
}

func NewllnessSymptomsSerializer(c *gin.Context, llnessSymptoms IllnessSymptomsModel) llnessSymptomsSerializer {
	return llnessSymptomsSerializer{
		c:                   c,
		IllnessSymptomsModel: llnessSymptoms,
	}
}

func NewllnessSymptomssSerializer(c *gin.Context, llnessSymptomss []IllnessSymptomsModel) llnessSymptomssSerializer {
	return llnessSymptomssSerializer{
		C:               c,
		llnessSymptomss: llnessSymptomss,
	}
}

func (s *llnessSymptomsSerializer) Response() llnessSymptomsResponse {
	return llnessSymptomsResponse{
		IllnessSymptomsModel: s.IllnessSymptomsModel,
	}
}

func (s *llnessSymptomssSerializer) Response() []llnessSymptomsResponse {
	response := make([]llnessSymptomsResponse, len(s.llnessSymptomss))
	for i, IllnessSymptomsModel := range s.llnessSymptomss {
		serializer := NewllnessSymptomsSerializer(s.C, IllnessSymptomsModel)
		response[i] = serializer.Response()
	}
	return response
}

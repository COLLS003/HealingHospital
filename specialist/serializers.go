package specialist

import (
	"github.com/gin-gonic/gin"
)

type SpecialistSerializer struct {
	c *gin.Context
	SpecialistModel
}

type SpecialistsSerializer struct {
	C           *gin.Context
	Specialists []SpecialistModel
}

type SpecialistResponse struct {
	SpecialistModel
}

func NewSpecialistSerializer(c *gin.Context, Specialist SpecialistModel) SpecialistSerializer {
	return SpecialistSerializer{
		c:               c,
		SpecialistModel: Specialist,
	}
}

func NewSpecialistsSerializer(c *gin.Context, Specialists []SpecialistModel) SpecialistsSerializer {
	return SpecialistsSerializer{
		C:           c,
		Specialists: Specialists,
	}
}

func (s *SpecialistSerializer) Response() SpecialistResponse {
	return SpecialistResponse{
		SpecialistModel: s.SpecialistModel,
	}
}

func (s *SpecialistsSerializer) Response() []SpecialistResponse {
	response := make([]SpecialistResponse, len(s.Specialists))
	for i, SpecialistModel := range s.Specialists {
		serializer := NewSpecialistSerializer(s.C, SpecialistModel)
		response[i] = serializer.Response()
	}
	return response
}

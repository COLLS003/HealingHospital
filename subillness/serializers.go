package subillness

import (
	"github.com/gin-gonic/gin"
)

type SubIllnessSerializer struct {
	c *gin.Context
	SubIllnessModel
}

type SubIllnesssSerializer struct {
	C           *gin.Context
	SubIllnesss []SubIllnessModel
}

type SubIllnessResponse struct {
	SubIllnessModel
}

func NewSubIllnessSerializer(c *gin.Context, SubIllness SubIllnessModel) SubIllnessSerializer {
	return SubIllnessSerializer{
		c:               c,
		SubIllnessModel: SubIllness,
	}
}

func NewSubIllnesssSerializer(c *gin.Context, SubIllnesss []SubIllnessModel) SubIllnesssSerializer {
	return SubIllnesssSerializer{
		C:           c,
		SubIllnesss: SubIllnesss,
	}
}

func (s *SubIllnessSerializer) Response() SubIllnessResponse {
	return SubIllnessResponse{
		SubIllnessModel: s.SubIllnessModel,
	}
}

func (s *SubIllnesssSerializer) Response() []SubIllnessResponse {
	response := make([]SubIllnessResponse, len(s.SubIllnesss))
	for i, SubIllnessModel := range s.SubIllnesss {
		serializer := NewSubIllnessSerializer(s.C, SubIllnessModel)
		response[i] = serializer.Response()
	}
	return response
}

package illness

import (
	"github.com/gin-gonic/gin"
)

type IllnessSerializer struct {
	c *gin.Context
	IllnessModel
}

type IllnesssSerializer struct {
	C        *gin.Context
	Illnesss []IllnessModel
}

type IllnessResponse struct {
	IllnessModel
}

func NewIllnessSerializer(c *gin.Context, Illness IllnessModel) IllnessSerializer {
	return IllnessSerializer{
		c:            c,
		IllnessModel: Illness,
	}
}

func NewIllnesssSerializer(c *gin.Context, Illnesss []IllnessModel) IllnesssSerializer {
	return IllnesssSerializer{
		C:        c,
		Illnesss: Illnesss,
	}
}

func (s *IllnessSerializer) Response() IllnessResponse {
	return IllnessResponse{
		IllnessModel: s.IllnessModel,
	}
}

func (s *IllnesssSerializer) Response() []IllnessResponse {
	response := make([]IllnessResponse, len(s.Illnesss))
	for i, IllnessModel := range s.Illnesss {
		serializer := NewIllnessSerializer(s.C, IllnessModel)
		response[i] = serializer.Response()
	}
	return response
}

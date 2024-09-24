package gorgonnx

import (
	"errors"
	"github.com/owulveryck/onnx-go"
	"gorgonia.org/gorgonia"
	"gorgonia.org/tensor"
)

type constantOfShape struct {
	value *tensor.Dense
}

func init() {
	register("ConstantOfShape", func() operator { return new(constantOfShape) })
}

func (c *constantOfShape) apply(graph *Graph, nodes ...*Node) error {
	if len(nodes) != 1 {
		return errors.New("wrong number of input nodes")
	}

	// 获取输入节点
	inputNode := nodes[0]

	//// 获取输入节点的子节点
	//children := getOrderedChildren(graph.g, inputNode)
	//// 检查输入节点的数量
	//err := checkCondition(children, 1)
	//if err != nil {
	//	return err
	//}

	s := inputNode.gorgoniaNode.Shape()

	//// 获取输入张量，即形状张量
	inputTensor := inputNode.t

	//// 创建具有指定形状的张量，并用填充值填充
	//newTensor, err := createFilledTensor(shape, fillValue)
	//if err != nil {
	//	return err
	//}

	// Empty dimensions in a tensor are not supported
	for i := range s {
		if s[i] <= 0 {
			return errors.New("empty dimensions are not allowed")
		}
	}

	t := tensor.New(tensor.WithShape(s...), tensor.Of(inputTensor.Dtype()))

	t, err := t.AddScalar(c.value, true)
	if err != nil {
		return err
	}

	// 设置输出节点
	inputNode.gorgoniaNode = gorgonia.NewConstant(t)

	return nil
}

func (c *constantOfShape) init(o onnx.Operation) error {
	attributes := o.Attributes
	if len(attributes) > 1 {
		return errors.New("only one attribute is allowed")
	}

	if len(attributes) == 1 {
		attr, ok := attributes["value"]
		if ok {
			c.value = tensor.New(tensor.WithBacking(attr))
			if c.value.Len() != 1 {
				return errors.New("value attribute must be a scalar")
			}
		} else {
			return errors.New("value attribute is missing")
		}
	} else {
		c.value = tensor.New(tensor.FromScalar(float32(0.0)))
	}
	return nil
}

// 获取填充值
//func getValue(o interface{}) *tensor.Dense {
//	// 默认填充值为 0.0
//	var fillValue float32 = 0.0
//	if attrs.Dtype() == tensor.Float32 {
//		fillValue = attrs.Data().(float32)
//	} else if attr.HasInt() {
//		fillValue = float32(attr.GetInt())
//	} else if attr.HasString() {
//		// 如果是字符串，尝试将其转换为 float32
//		value, err := strconv.ParseFloat(string(attr.GetString()), 32)
//		if err == nil {
//			fillValue = float32(value)
//		}
//	}
//
//	value := tensor.New(tensor.WithBacking(o))
//
//	return fillValue
//}

func IfScalarToSlice(value any) any {
	switch data := value.(type) {
	case int8:
		return []int8{data}
	case int16:
		return []int16{data}
	case int32:
		return []int32{data}
	case int64:
		return []int64{data}
	case int:
		return []int{data}
	case float32:
		return []float32{data}
	case float64:
		return []float64{data}
	case complex64:
		return []complex64{data}
	case complex128:
		return []complex128{data}
	default:
		return value
	}
}

// AnyToIntSlice casts the data of a node to an int list. This will only
// be done if the data is of some sort of int type.
func AnyToIntSlice(value interface{}) ([]int, error) {
	var res []int

	switch data := value.(type) {
	case []int8:
		for _, value := range data {
			res = append(res, int(value))
		}

		return res, nil
	case []int16:
		for _, value := range data {
			res = append(res, int(value))
		}

		return res, nil
	case []int32:
		for _, value := range data {
			res = append(res, int(value))
		}

		return res, nil
	case []int64:
		for _, value := range data {
			res = append(res, int(value))
		}

		return res, nil
	default:
		return nil, errors.New("constant of shape error")
	}
}

package gorgonnx

//
//import (
//	"errors"
//	"github.com/owulveryck/onnx-go"
//	"gorgonia.org/gorgonia"
//	"gorgonia.org/tensor"
//)
//
//func init() {
//	register("Cast", func() operator { return new(cast) })
//}
//
//type cast struct{}
//
//func (*cast) apply(graph *Graph, nodes ...*Node) error {
//	if len(nodes) != 1 {
//		return errors.New("wrong number of input nodes")
//	}
//
//	// 获取输入节点
//	inputNode := nodes[0]
//
//	// 获取输入节点的子节点
//	children := getOrderedChildren(graph.g, inputNode)
//
//	// 检查输入节点的数量
//	err := checkCondition(children, 1)
//	if err != nil {
//		return err
//	}
//
//	// 获取输入张量
//	inputTensor := children[0].gorgoniaNode.Value().(*tensor.Tensor)
//
//	// 获取目标数据类型
//	targetType := getTargetDataType(inputNode.Attributes["to"])
//
//	// 将输入张量转换为目标数据类型
//	castedTensor, err := castTensor(inputTensor, targetType)
//	if err != nil {
//		return err
//	}
//
//	// 设置输出节点
//	inputNode.gorgoniaNode = gorgonia.NewConstant(castedTensor)
//
//	return nil
//}
//
//func (*cast) init(op onnx.Operation) error {
//	return nil
//}
//
//// 获取目标数据类型
//func getTargetDataType(attr onnx.AttributeProto) tensor.DType {
//	switch attr.GetInt() {
//	case int(onnx.INT32):
//		return tensor.Int32
//	case int(onnx.INT64):
//		return tensor.Int64
//	case int(onnx.FLOAT):
//		return tensor.Float32
//	case int(onnx.DOUBLE):
//		return tensor.Float64
//	default:
//		panic("unsupported data type")
//	}
//}
//
//// 将张量转换为目标数据类型
//func castTensor(t *tensor.Tensor, targetType tensor.DType) (*tensor.Tensor, error) {
//	switch targetType {
//	case tensor.Int32:
//		return t.AsType(targetType, false)
//	case tensor.Int64:
//		return t.AsType(targetType, false)
//	case tensor.Float32:
//		return t.AsType(targetType, false)
//	case tensor.Float64:
//		return t.AsType(targetType, false)
//	default:
//		return nil, errors.New("unsupported data type for casting")
//	}
//}

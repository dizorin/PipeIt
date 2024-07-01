package pipe

import (
	"encoding/gob"
	"fmt"
	imgui "github.com/AllenDang/cimgui-go"
	g "github.com/AllenDang/giu"
	"github.com/PaesslerAG/gval"
	"strings"
)

type ScriptPipe struct {
	Script string
	Err    error
}

func init() {
	gob.Register(&ScriptPipe{})
}

func NewScriptPipe() Pipe {
	return &ScriptPipe{}
}

func (m *ScriptPipe) GetName() string {
	return "Sc"
}

func (m *ScriptPipe) GetTip() string {
	return fmt.Sprintf("Run script for each string of input string array. Available variables: [id, value]")
}

func (m *ScriptPipe) GetInputType() DataType {
	return DataTypeStringArray
}

func (m *ScriptPipe) GetOutputType() DataType {
	return DataTypeStringArray
}

func (m *ScriptPipe) GetConfigUI(changed func()) g.Layout {
	return g.Layout{
		g.Style().
			SetColor(g.StyleColorText, g.Vec4ToRGBA(imgui.NewVec4(0.8, 0.4, 0.4, 1.))).
			To(g.Condition(m.Err != nil, g.Custom(func() { g.Label(m.Err.Error()).Build() }), nil)),
		g.InputText(&(m.Script)).
			Flags(g.InputTextFlagsCallbackCompletion | g.InputTextFlagsCallbackHistory).
			Label("Script").Size(400).OnChange(changed),
	}
}

func (m *ScriptPipe) Process(data interface{}) interface{} {
	if strs, ok := data.([]string); ok {
		var parameter = make(map[string]interface{})
		var result []string

		m.Err = nil
		for i, s := range strs {

			parameter["id"] = i
			parameter["value"] = s

			value, err := gval.Evaluate(m.Script, parameter)
			if err != nil && strings.Contains(err.Error(), "EOF") {
				continue
			} else if err != nil {
				result = append(result, fmt.Sprintf("{Error: %v}", err))
				m.Err = err
				return result
			}
			result = append(result, fmt.Sprintf("%v", value))
		}

		return result
	}

	return []string{"Error: Script only accepts string array as input"}
}

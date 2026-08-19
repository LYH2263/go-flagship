package eval

import "context"

// RunAllContext 批量评估并尊重取消。
func RunAllContext(ctx context.Context, inputs []Input) ([]Result, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	out := make([]Result, 0, len(inputs))
	for _, in := range inputs {
		if err := ctx.Err(); err != nil {
			return nil, err
		}
		out = append(out, Run(in))
	}
	return out, nil
}

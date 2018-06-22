package vector

// Vector is 2D vector
type Vector struct {
	X, Y float64
}

// Add is v + o
func (v Vector) Add(o Vector) Vector {
	return Vector{v.X + o.X, v.Y + o.Y}
}

// Sub is v - o
func (v Vector) Sub(o Vector) Vector {
	return Vector{v.X - o.X, v.Y - o.Y}
}

// Dot is v * Scalar
func (v Vector) Dot(f float64) Vector {
	return Vector{v.X * f, v.Y * f}
}

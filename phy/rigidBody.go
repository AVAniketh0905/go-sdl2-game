package phy

const UNIT_MASS = 1.0
const GRAVITY = 9.8
const DAMPING_COEFF = 0.9

const MAX_JUMP_HEIGHT = 20
const MAX_JUMP_TIME = 15

const MAX_VELOCITY = 100

type RigidBody struct {
	mass float64

	position     *Vector
	displacement *Vector
	velocity     *Vector
	acceleration *Vector

	forces   []Vector
	gravity  *Vector
	friction *Vector
}

func NewRigidBody(transform *Transform) *RigidBody {
	g := 10 + 0*float64(2*MAX_JUMP_HEIGHT)/(MAX_JUMP_TIME*MAX_JUMP_TIME)

	return &RigidBody{
		mass:         UNIT_MASS,
		position:     &Vector{X: transform.X, Y: transform.Y},
		displacement: &Vector{X: 0, Y: 0},
		velocity:     &Vector{X: 0, Y: 0},
		acceleration: &Vector{X: 0, Y: 0},
		friction:     &Vector{X: 0, Y: 0},
		gravity:      &Vector{X: 0, Y: -g},
		forces:       []Vector{},
	}
}

func (rb *RigidBody) GetGravity() *Vector {
	return rb.gravity
}

func (rb *RigidBody) SetPosition(pos *Vector) {
	rb.position = pos
}

func (rb *RigidBody) AddVelocity(velocity Vector) {
	rb.velocity.Add(&velocity)
	rb.velocity.Clamp(-MAX_VELOCITY, MAX_VELOCITY)
}

func (rb *RigidBody) AddForce(force Vector) {
	rb.forces = append(rb.forces, force)
}

func (rb *RigidBody) AddGravity() {
	rb.AddForce(*rb.gravity)
}

func (rb *RigidBody) ApplyFriction(dt float64) {
	friction := 1 + DAMPING_COEFF*dt
	rb.velocity.Div(friction)
}

func (rb *RigidBody) ApplyForces(dt float64) {
	rb.acceleration = &Vector{X: 0, Y: 0}
	rb.AddGravity()

	for _, force := range rb.forces {
		force.Div(rb.mass)
		rb.acceleration.Add(&force)
	}
}

func (rb *RigidBody) UnsetForces() {
	rb.forces = []Vector{}
	rb.acceleration = &Vector{X: 0, Y: 0}
	rb.velocity = &Vector{X: 0, Y: 0}
}

// TODO
func (rb *RigidBody) Update(dt uint64) {
	rb.displacement = rb.position.Copy()

	rb.ApplyForces(float64(dt))
	// rb.acceleration = rb.gravity.Copy()

	rb.velocity.X += rb.acceleration.X * float64(dt)
	rb.velocity.Y += rb.acceleration.Y * float64(dt)

	rb.velocity.Div(1 + DAMPING_COEFF*float64(dt))
	rb.velocity.Clamp(-MAX_VELOCITY, MAX_VELOCITY)

	// acc_copied := rb.acceleration.Copy()
	// acc_copied.Mult(float64(dt))
	// rb.velocity.Add(acc_copied)

	// rb.ApplyFriction(float64(dt))

	rb.position.X += rb.velocity.X * float64(dt)
	rb.position.Y += rb.velocity.Y * float64(dt)

	// vel_copied := rb.velocity.Copy()
	// vel_copied.Mult(float64(dt))
	// rb.position.Add(vel_copied)

	rb.displacement.Sub(rb.position)
}

// Getters
func (rb *RigidBody) GetMass() float64 {
	return rb.mass
}

func (rb *RigidBody) GetPosition() *Vector {
	return rb.position
}

func (rb *RigidBody) GetDisplacement() *Vector {
	return rb.displacement
}

func (rb *RigidBody) GetVelocity() *Vector {
	return rb.velocity
}

func (rb *RigidBody) GetAcceleration() *Vector {
	return rb.acceleration
}

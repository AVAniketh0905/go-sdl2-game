package phy

const UNIT_MASS = 1.0
const GRAVITY = 9.8
const FRICTION_COEFF = 0.1

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
	return &RigidBody{
		mass:         UNIT_MASS,
		position:     &Vector{X: transform.X, Y: transform.Y},
		displacement: &Vector{X: 0, Y: 0},
		velocity:     &Vector{X: 0, Y: 0},
		acceleration: &Vector{X: 0, Y: 0},
		friction:     &Vector{X: 0, Y: 0},
		gravity:      &Vector{X: 0, Y: -GRAVITY},
		forces:       []Vector{},
	}
}

func (rb *RigidBody) SetPosition(pos *Vector) {
	rb.position = pos
}

func (rb *RigidBody) AddVelocity(velocity Vector) {
	rb.velocity.Add(&velocity)
}

func (rb *RigidBody) AddForce(force Vector) {
	rb.forces = append(rb.forces, force)
}

func (rb *RigidBody) AddGravity() {
	rb.AddForce(*rb.gravity)
}

func (rb *RigidBody) ApplyFriction() {
	friction := rb.velocity.Copy()
	friction.Normalize()
	friction.Mult(-1)
	friction.Mult(FRICTION_COEFF)
	rb.AddForce(*friction)
}

func (rb *RigidBody) ApplyForces() {
	rb.acceleration = &Vector{X: 0, Y: 0}
	rb.AddGravity()

	for _, force := range rb.forces {
		force.Div(rb.mass)
		rb.acceleration.Add(&force)
	}

	rb.ApplyFriction()
}

func (rb *RigidBody) UnsetForces() {
	rb.forces = []Vector{}
	rb.acceleration = &Vector{X: 0, Y: 0}
	rb.velocity = &Vector{X: 0, Y: 0}
	rb.position = &Vector{X: 0, Y: 0}
}

// TODO
func (rb *RigidBody) Update(dt float64) {
	rb.displacement = rb.position.Copy()

	rb.ApplyForces()

	acc_copied := rb.acceleration.Copy()
	acc_copied.Mult(dt)
	rb.velocity.Add(acc_copied)

	vel_copied := rb.velocity.Copy()
	vel_copied.Mult(dt)
	rb.position.Add(vel_copied)

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

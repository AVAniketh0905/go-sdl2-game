package phy

const UNIT_MASS = 1.0
const GRAVITY = 0.098
const FRICTION_COEFF = 0.1

type RigidBody struct {
	mass float64

	position     *Vector
	velocity     *Vector
	acceleration *Vector

	forces   []Vector
	friction *Vector
}

func NewRigidBody(position *Vector) *RigidBody {
	return &RigidBody{
		mass:         UNIT_MASS,
		position:     position,
		velocity:     &Vector{X: 0, Y: 0},
		acceleration: &Vector{X: 0, Y: 0},
		friction:     &Vector{X: 0, Y: 0},
		forces: []Vector{
			{X: 0, Y: GRAVITY * UNIT_MASS},
		},
	}
}

func (rb *RigidBody) AddForce(force Vector) {
	rb.forces = append(rb.forces, force)
}

func (rb *RigidBody) ApplyFriction() {
	friction := rb.velocity.Copy()
	friction.Normalize()
	friction.Mult(-1)
	friction.Mult(FRICTION_COEFF)
	rb.AddForce(*friction)
}

func (rb *RigidBody) ApplyForces() {
	for _, force := range rb.forces {
		force.Div(rb.mass)
		rb.acceleration.Add(&force)
	}
	rb.ApplyFriction()
}

func (rb *RigidBody) Update(dt float64) {
	// log.Println("Updating RigidBody", rb.position, rb.velocity, rb.acceleration)
	rb.ApplyForces()
	rb.acceleration.Mult(dt)
	rb.velocity.Add(rb.acceleration)
	rb.velocity.Mult(dt)
	rb.position.Add(rb.velocity)
}

// Getters
func (rb *RigidBody) GetMass() float64 {
	return rb.mass
}

func (rb *RigidBody) GetPosition() *Vector {
	return rb.position
}

func (rb *RigidBody) GetVelocity() *Vector {
	return rb.velocity
}

func (rb *RigidBody) GetAcceleration() *Vector {
	return rb.acceleration
}

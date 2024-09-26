package fr24

type Fr24Error struct {
	Err string
}

func (e Fr24Error) Error() string {
	return e.Err
}

func NewFr24Error(err error) Fr24Error {
	/**
	* Factory function to make a new Fr24Error object.
	 */
	return Fr24Error{Err: err.Error()}
}

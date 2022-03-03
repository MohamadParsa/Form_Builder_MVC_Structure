package models

/*
br struct contains any property of a br element that we want to change dynamically.
any of these properties used in the corresponding template. (you can see this file: "viewmodels/html objects/br")
*/
type br struct {
	ID string
}

/*
getData is a method of br type and get HTML element id and return required data to create a br in the page.
this method returns a br type but in function declaration uses the interface. it means this method can return any
other type so we can define this method for other elements like textbox,...
*/
func (br) getData(id string) interface{} {
	/*
		in this step, we should connect to a database or any other storage and send "id" then get br information.
		fetching data don't implement so just a piece of constant information has been returned.
		you need to develop this section as your logic.
	*/
	return br{ID: id}
}

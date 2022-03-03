package models

/*
button struct contains any property of a button element that we want to change dynamically.
any of these properties used in the corresponding template. (you can see this file: "viewmodels/html objects/button")
*/
type button struct {
	ID      string
	Caption string
}

/*
getData is a method of button type and get HTML element id and return required data to create a button in the page.
this method returns a button type but in function declaration uses the interface. it means this method can return any
other type so we can define this method for other elements like textbox,...
*/
func (button) getData(id string) interface{} {
	/*
		in this step, we should connect to a database or any other storage and send "id" then get button information.
		fetching data don't implement so just a piece of constant information has been returned.
		you need to develop this section as your logic.
	*/
	return button{ID: id, Caption: "Login"}
}

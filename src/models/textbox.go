package models

/*
textbox struct contains any property of a text box element that we want to change dynamically.
any of these properties used in the corresponding template. (you can see this file: "viewmodels/html objects/textbox")
*/
type textbox struct {
	ID          string
	Lable       string
	PlaceHolder string
	Value       string
}

/*
getData is a method of textbox type and get HTML element id and return required data to create a text box in the page.
this method returns a textbox type but in function declaration uses the interface. it means this method can return any
other type so we can define this method for other elements like button,...
*/
func (textbox) getData(id string) interface{} {
	/*
		in this step, we should connect to a database or any other storage and send "id" then get text box information.
		fetching data don't implement so just a piece of constant information has been returned.
		you need to develop this section as your logic.
	*/
	return textbox{ID: id, Lable: "User Name", PlaceHolder: "Emaile, Phone Number,..."}

}

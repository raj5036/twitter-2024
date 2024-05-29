import axios from "axios"

export const getUser = (params: any) => {
	console.log(params)
	return Promise.resolve(
		axios.get('')
	)
		.then(response => {

		})
		.catch(error => {
			
		})
}
import { GetECS, GetRDS } from '@/http/resources';
export const GetDelay = () => {

	// let good = 0
	let seven = 0
	let thirty = 0
	let expired = 0

	GetECS().then(res => {
		// console.log(res);
		let content = JSON.parse(res.data.msg)

		for (let i = 0; i < content.length; i++) {
			// if (content[i].status == '0') {
			// 	good++
			// }
			if (content[i].status == '2') {
				seven++
				continue
			}
			if (content[i].status == '1') {
				thirty++
				continue
			}
			if (content[i].status == '3') {
				expired++
			}
		}

	})
	GetRDS().then((res) => {
		// console.log(res);
		let content = JSON.parse(res.data.msg)
		for (let i = 0; i < content.length; i++) {
			if (content[i].status == '2') {
				seven++
				continue
			}
			if (content[i].status == '1') {
				thirty++
				continue
			}
			if (content[i].status == '3') {
				expired++
			}
		}
	})

	return [{
		// good,
		seven,
		thirty,
		expired,
	}]
}
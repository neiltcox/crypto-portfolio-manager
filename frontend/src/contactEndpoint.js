import Cookies from 'universal-cookie';
const cookies = new Cookies();

export default async function contactEndpoint(method, resource, query = {}, body = {}){
	method = method.toUpperCase();

	if(!['GET', 'POST', 'DELETE'].includes(method)) throw 'Invalid request method';

	let raw_response = null;

	try {
		let options = {
			method: method,
			cache: 'no-cache',
			headers: {
				'Content-Type': 'application/json'
			}
		};

		/*if (cookies.get('magscan_login_token')) {
			options.headers['X-Login-Token'] = cookies.get('magscan_login_token');
		}*/

		if(method === 'POST') options.body = JSON.stringify(body);

		raw_response = await fetch(
			"/api/v1/"+resource+(Object.keys(query).length > 0 ? ("?"+(new URLSearchParams(query).toString())) : ""),
			options
		);
	}catch(err){
		throw err.message;
	}

	let json = await raw_response.json();
	if(raw_response.status != 200) throw json.Message+" (status code "+raw_response.status+")";

	if(!json) throw "Server response could not be parsed";
	
	return json;
}
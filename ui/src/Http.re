let get = (url) =>
	Js.Promise.(
		Fetch.fetch(url)
		|> then_(Fetch.Response.text)
		|> then_(text => print_endline(text) |> resolve)
	);

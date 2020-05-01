let baseUrl = "http://localhost:8080";

type account = {
  id: string,
  name: string,
  createdAt: string,
  updatedAt: string,
};

type accounts = array(account);

module Decode = {
  let account = (json): account =>
    Json.Decode.{
      id: json |> field("id", string),
      name: json |> field("name", string),
      createdAt: json |> field("createdAt", string),
      updatedAt: json |> field("updatedAt", string),
    };

  let accounts = (json): array(account) =>
    Json.Decode.(json |> array(account));
};

let listAccounts = callback =>
  Js.Promise.(
    Fetch.fetch(baseUrl ++ "/account")
    |> then_(Fetch.Response.json)
    |> then_(json =>
         json
         |> Decode.accounts
         |> (
           accounts => {
             callback(accounts);
             resolve();
           }
         )
       )
    |> ignore
  );

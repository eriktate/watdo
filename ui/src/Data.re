let baseUrl = "http://localhost:8080";

type account = {
  id: string,
  name: string,
  createdAt: string,
  updatedAt: string,
};

type accounts = array(account);

let emptyAccount = {id: "", name: "", createdAt: "", updatedAt: ""};

let withId = (id, fieldList) => {
  if (id != "") {
    [("id", Json.Encode.string(id))] @ fieldList
  } else {
    fieldList
  }
}

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

  let id = (json): string => Json.Decode.string(json);
};

module Encode = {
  let account = account => {
    let fieldList = [("name", Json.Encode.string(account.name))];

    Json.Encode.(object_(withId(account.id, fieldList)));
  };
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

let saveAccount = (account, callback) =>
  Js.Promise.(
    Fetch.fetchWithInit(
      baseUrl ++ "/account",
      Fetch.RequestInit.make(
        ~method_=Fetch.Post,
        ~body=Fetch.BodyInit.make(Json.stringify(Encode.account(account))),
        (),
      ),
    )
    |> then_(Fetch.Response.json)
    |> then_(json =>
         json
         |> Decode.id
         |> (
           id => {
             callback(id);
             resolve();
           }
         )
       )
    |> ignore
  );

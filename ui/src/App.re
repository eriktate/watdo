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

let listAccounts = (url, callback) =>
  Js.Promise.(
    Fetch.fetch(url)
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

type action =
  | Loaded(accounts)
  | Loading;

type state = {
  accounts,
  loading: bool,
};

let initialState = {accounts: [||], loading: false};

[@react.component]
let make = () => {
  let (state, dispatch) =
    React.useReducer(
      (state, action) =>
        switch (action) {
        | Loading => {...state, loading: true}
        | Loaded(accounts) => {...state, accounts}
        },
      initialState,
    );

  React.useEffect0(() => {
    listAccounts("http://localhost:8080/account", payload =>
      dispatch(Loaded(payload))
    )
    |> ignore;
    None;
  });

  <ul>
    {state.accounts
     |> Array.map(account =>
          <li key={account.id}>
            {ReasonReact.string(account.id ++ ": " ++ account.name)}
          </li>
        )
     |> React.array}
  </ul>;
};

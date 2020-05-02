open Data;

type state = {
  accounts,
  selectedAccount: option(string),
  loadingAccounts: bool,
};

type action =
  | LoadingAccounts
  | LoadedAccounts(accounts)
  | SelectAccount(string);

let initialState = {
  accounts: [||],
  selectedAccount: None,
  loadingAccounts: false,
};

let refreshAccounts = dispatch =>
  listAccounts(payload => dispatch(LoadedAccounts(payload)));

let findAccountById = (accounts, id) =>
  Array.to_list(accounts) |> List.find(account => account.id == id);

[@react.component]
let make = () => {
  let (state, dispatch) =
    React.useReducer(
      (state, action) =>
        switch (action) {
        | LoadingAccounts => {...state, loadingAccounts: true}
        | LoadedAccounts(accounts) => {
            ...state,
            loadingAccounts: false,
            accounts,
          }
        | SelectAccount(id) => {...state, selectedAccount: Some(id)}
        },
      initialState,
    );

  let refreshAccounts = () => refreshAccounts(dispatch) |> ignore;
  let selectAccount = id => {
    Js.log2("Selecting account: ", id);
    dispatch(SelectAccount(id)) |> ignore;
  };

  let account =
    switch (state.selectedAccount) {
    | Some(id) => Some(findAccountById(state.accounts, id))
    | None => None
    };

  <div>
    <AccountEdit account refreshAccounts />
    <Accounts accounts={state.accounts} refreshAccounts selectAccount />
  </div>;
};

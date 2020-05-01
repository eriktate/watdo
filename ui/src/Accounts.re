open Data;

type state = {
  accounts,
  loading: bool,
};

type action =
  | Loaded(accounts)
  | Loading;

let initialState = {accounts: [||], loading: false};

let listHelper = (dispatch, _event) => {
  dispatch(Loading);
  listAccounts(payload => dispatch(Loaded(payload)));
};

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
    listHelper(dispatch, ());
    None;
  });

  <div>
    <button onClick={listHelper(dispatch)}>
      {ReasonReact.string("Refresh")}
    </button>
    <ul>
      {state.accounts
       |> Array.map(account =>
            <li key={account.id}> {ReasonReact.string(account.name)} </li>
          )
       |> React.array}
    </ul>
  </div>;
};

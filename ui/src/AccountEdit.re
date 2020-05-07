type state = {
  saving: bool,
  account: Data.account,
};

type action =
  | SetAccount(Data.account)
  | SetName(string)
  | Saving
  | Saved
  | Clear;

let initialState = account => {
  let initAccount =
    switch (account) {
    | Some(acc) => acc
    | None => Data.emptyAccount
    };
  {saving: false, account: initAccount};
};

let saveAccount = (dispatch, account, _event) => {
  dispatch(Saving);
  Data.saveAccount(
    account,
    payload => {
      Js.log2("Saved account: ", payload);
      dispatch(Saved);
    },
  );
};

[@react.component]
let make = (~account, ~refreshAccounts) => {
  let (state, dispatch) =
    React.useReducer(
      (state, action) =>
        switch (action) {
        | SetAccount(account) => {...state, account}
        | SetName(name) => {
            ...state,
            account: {
              ...state.account,
              name,
            },
          }
        | Saving => {...state, saving: true}
        | Saved =>
          refreshAccounts();
          {saving: false, account: Data.emptyAccount};
        | Clear => {...state, account: Data.emptyAccount}
        },
      initialState(account),
    );

  React.useEffect1(
    () => {
      switch (account) {
      | Some(account) =>
        dispatch(SetAccount(account));
        None;
      | None => None
      }
    },
    [|account|],
  );

  let saveAccount = saveAccount(dispatch);
  let setName = event => {
    let name = ReactEvent.Form.target(event)##value;
    dispatch(SetName(name));
  };
  let clear = _event => dispatch(Clear);

  <form className="tile">
    {if (state.account.id == "") {
       ReasonReact.string("Create new account");
     } else {
       ReasonReact.string(
         "Editing " ++ state.account.name ++ ": " ++ state.account.id,
       );
     }}
    <div className="fieldSet">
      <label>
        {ReasonReact.string("Account Name:")}
        <input
          value={state.account.name}
          type_="input"
          name="accountName"
          onChange=setName
        />
      </label>
    </div>
    <button type_="button" onClick={saveAccount(state.account)} className="btn-primary">
      {ReasonReact.string("Save")}
    </button>
    <button type_="button" onClick=clear className="btn-secondary">
      {ReasonReact.string("Clear")}
    </button>
  </form>;
};

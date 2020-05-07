type state = {selecting: bool};

type action =
  | Open
  | Close;

let initialState = {selecting: false};

let findItem = (items: Data.listItems, id) =>
  try(
    List.find((item: Data.listItem) => item.id == id, Array.to_list(items))
  ) {
  | Not_found => {id: "", name: ""}
  };

[@react.component]
let make = (~items, ~selectedId, ~select) => {
  let (state, dispatch) =
    React.useReducer(
      (_state, action) => {
        switch (action) {
        | Open => {selecting: true}
        | Close => {selecting: false}
        }
      },
      initialState,
    );

  let select = (id, _event) => {
    dispatch(Close);
    select(id);
  };

  let ddOpen = _event => dispatch(Open);

  <div className="dropdown">
    {if (state.selecting) {
       items
       |> Array.map((item: Data.listItem) =>
            <a key={item.id} href="#" onClick={select(item.id)}>
              {ReasonReact.string(item.name)}
            </a>
          )
       |> React.array;
     } else {
       <a href="#" onClick=ddOpen>
         {ReasonReact.string(findItem(items, selectedId).name)}
       </a>;
     }}
  </div>;
};

[@react.component]
let make = () => {
  <nav>
    <div id="banner">
    </div>
    <ul>
      <li>
        <a href="/">{ReasonReact.string("Dashboard")}</a>
      </li>
      <li>
        <a href="/tasks">{ReasonReact.string("Tasks")}</a>
      </li>
    </ul>
  </nav>;
};

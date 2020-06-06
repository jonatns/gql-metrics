import React from 'react';
import logo from './logo.svg';
import './App.css';

function App() {
  const [queries, setQueries] = React.useState([]);
  React.useEffect(() => {
    const fetchQueries = async () => {
      const resp = await fetch('/api/queries');
      const data = await resp.json();
      setQueries(data);
    };

    fetchQueries();
  }, []);

  if (!queries.length) {
    return <>Loading...</>;
  }

  return (
    <div className="App">
      {queries.map((query) => (
        <div>{query.name}</div>
      ))}
    </div>
  );
}

export default App;

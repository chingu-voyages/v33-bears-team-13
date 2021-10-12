import Head from 'next/head';
import React, { useState, useEffect } from 'react';
import axios from 'axios';

export default function Home() {
  const [result, setResult] = useState('');
  const [city, setCity] = useState('london');
  const [postresult, setPostresult] = useState('Ready to save');
  const [summarylist, setSummarylist] = useState([
    'New York',
    'Berlin',
    'London',
    'Paris',
  ]);

  useEffect(() => {
    setResult("Let's get weather");
  });

  const handleClick = () => {
    axios
      .get(`http://localhost:8080/weather/` + city)
      .then(function (response) {
        //var res = JSON.parse(response.data);
        var res = response.data;
        console.log(res);

        setResult(res);
      })
      .catch(function (error) {
        setResult('error!');
      });
  };

  useEffect(() => {
    setResult(result);
  }, [result]);

  const handleChange = (e) => {
    e.preventDefault();
    console.log(e.target.value);
    setCity(e.target.value);
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    handleClick(city);
    console.log(city);
  };

  const handleSave = (input) => {
    // console.log(input);

    axios
      .post(`http://localhost:8080/summaries`, `"${result}"`)
      .then(function (response) {
        //res = JSON.parse(response.data);
        console.log(response);

        //let items = res.items[0];

        setPostresult('Saved!!!');
      })
      .catch(function (error) {
        setPostresult('error!');
      });
  };

  const getList = () => {
    axios
      .get(`http://localhost:8080/summaries`)
      .then(function (response) {
        //res = JSON.parse(response.data);
        console.log(response.data);

        //let items = res.items[0];

        setSummarylist(JSON.stringify(response.data));
      })
      .catch(function (error) {
        setSummarylist('error!');
      });
  };

  return (
    <div>
      <form onSubmit={handleSubmit}>
        <label>
          City==============>
          <input type="text" onChange={handleChange} />
        </label>
        <input type="submit" value="Submit" />
      </form>
      <div>Result=============>{result}</div>
      <button onClick={() => handleSave(result)}>Save Button Here</button>
      <div>Save Result=========>{postresult}</div>
      <button onClick={() => getList()}>Get list Button Here</button>
      <div>{Record(summarylist)}</div>
    </div>
  );
}

function Record(record) {
  return (
    <div className="pl-6 pr-6">
      <h2 className="border-b-2">Summaries record</h2>
      <ul>
        {record.map((summary) => (
          <li>{summary}</li>
        ))}
      </ul>
    </div>
  );
}

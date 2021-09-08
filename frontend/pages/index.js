import Head from 'next/head';
import React, { useState, useEffect } from 'react';
import axios from 'axios';

export default function Home() {
  const [result, setResult] = useState('');
  const [city, setCity] = useState('london');

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

        // let items = res[0];

        var renewedResult =
          'Current temperature in ' +
          res.name +
          ' is ' +
          res.main.temp +
          '°C .................Conditions are currently ' +
          res.weather[0].description;

        setResult(renewedResult);
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

  return (
    <div>
      <form onSubmit={handleSubmit}>
        <label>
          City::::::::::
          <input type="text" onChange={handleChange} />
        </label>
        <input type="submit" value="Submit" />
      </form>
      <div>{result}</div>
    </div>
  );
}

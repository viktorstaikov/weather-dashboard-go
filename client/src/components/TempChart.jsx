import moment from 'moment';
import React from 'react';
import { Area, AreaChart, Label, ResponsiveContainer, Tooltip, XAxis, YAxis } from 'recharts';
import Title from './Title';

function genLabel(value) {
  const m = moment(value);

  const startOfDay = moment(value);
  const today = moment();

  let s = value;

  if (startOfDay.isSame(today, 'day')) {
    s = 'Today, ' + m.format('HH:mm');
  } else {
    s = m.format('ddd, HH:mm');
  }
  return s;
}

export default function TempChart(props) {
  const { series } = props;
  const mapped = !series
    ? []
    : series.map(item => {
        return {
          label: genLabel(item.timestamp),
          min: item.temp_min,
          max: item.temp_max,
        };
      });
  return (
    <React.Fragment>
      <Title>Temperature&deg;</Title>
      <ResponsiveContainer>
        <AreaChart
          data={mapped}
          margin={{
            top: 16,
            right: 16,
            bottom: 0,
            left: 24,
          }}
        >
          <defs>
            <linearGradient id="colorMin" x1="0" y1="0" x2="0" y2="1">
              <stop offset="5%" stopColor="blue" stopOpacity={0.8} />
              <stop offset="95%" stopColor="blue" stopOpacity={0} />
            </linearGradient>
            <linearGradient id="colorMax" x1="0" y1="0" x2="0" y2="1">
              <stop offset="5%" stopColor="orange" stopOpacity={0.8} />
              <stop offset="95%" stopColor="orange" stopOpacity={0} />
            </linearGradient>
          </defs>

          <Tooltip />
          <XAxis dataKey="label"></XAxis>
          <YAxis>
            <Label angle={270} position="left" style={{ textAnchor: 'middle' }}>
              Degrees Celsius
            </Label>
          </YAxis>

          <Area type="monotone" dataKey="max" stroke="red" fillOpacity={1} fill="url(#colorMax)" />
          <Area type="monotone" dataKey="min" stroke="blue" fillOpacity={1} fill="url(#colorMin)" />
        </AreaChart>
      </ResponsiveContainer>
    </React.Fragment>
  );
}

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

export default function DataChart(props) {
  const { series, title, yAxisLabel, color } = props;

  const mapped = !series
    ? []
    : series.map(item => {
        return {
          label: genLabel(item.timestamp),
          value: item.value,
        };
      });
  return (
    <React.Fragment>
      <Title>{title}</Title>
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
            <linearGradient id="grad" x1="0" y1="0" x2="0" y2="1">
              <stop offset="5%" stopColor={color} stopOpacity={0.8} />
              <stop offset="95%" stopColor={color} stopOpacity={0} />
            </linearGradient>
          </defs>

          <Tooltip />
          <XAxis dataKey="label"></XAxis>
          <YAxis>
            <Label angle={270} position="left" style={{ textAnchor: 'middle' }}>
              {yAxisLabel}
            </Label>
          </YAxis>

          <Area type="monotone" dataKey="value" stroke={color} fillOpacity={1} fill="url(#grad)" />
        </AreaChart>
      </ResponsiveContainer>
    </React.Fragment>
  );
}

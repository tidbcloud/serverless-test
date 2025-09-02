'use client';

import React from 'react';
import { BarChart, Bar, XAxis, YAxis, Tooltip, ResponsiveContainer, Cell } from 'recharts';

interface DataItem {
  plan: string;
  date: string;
  availability?: number;
  p99?: number;
  total?: number;
  failed_count?: number;
  success_count?: number;
}

interface Props {
  data: DataItem[];
  lastNDays: string[];
}

export default function ProbeChart({ data = [], lastNDays = [] }: Props) {
  const plans = ['starter', 'essential'];
  const planColors: Record<string, string> = {
    starter: '#4ade80',
    essential: '#60a5fa',
  };

  return (
    <div className="w-full">
      {plans.map((plan) => {
        const planDataMap = new Map(
          data.filter(d => d.plan === plan).map(d => [d.date, d])
        );
        
        const fullData = lastNDays.map(date => {
          const existingData = planDataMap.get(date);
          return {
            date,
            plan,
            isEmpty: !existingData,
            availability: existingData?.availability,
            p99: existingData?.p99,
            count: existingData?.total,
            failed_count: existingData?.failed_count,
            succeed_count: existingData?.success_count
          };
        });
        
        const validData = fullData.filter(d => !d.isEmpty);
        const totalRequests = validData.reduce((sum, d) => sum + +(d.count || 0), 0);
        // don't know why sum + (d.failed_count || 0) will concatenate as string, for example: failed_count 3 and 4 results in "034". 
        // so we use +(d.failed_count || 0) to force it to be number. The + means convert to number.
        const totalFails = validData.reduce((sum, d) => sum + +(d.failed_count || 0), 0);

              
        const overallAvailability = totalRequests > 0
          ? Number(((totalRequests - totalFails) / totalRequests * 100).toFixed(2))
          : null;
        const isHealthy = overallAvailability != null && overallAvailability >= 99;

        return (
          <div key={plan} className={`w-full bg-white rounded-2xl shadow-sm border border-gray-200 flex flex-col px-6 py-6 sm:px-8 sm:py-8 ${plan === 'essential' ? 'mt-16' : ''}`}>
            <div className="flex items-center justify-between w-full mb-3">
              <h2 className="text-2xl font-normal text-gray-900">
                {plan.charAt(0).toUpperCase() + plan.slice(1)}
              </h2>
              <div className="flex items-center gap-2">
                <div className={`w-2 h-2 rounded-full ${isHealthy ? 'bg-green-500' : 'bg-red-500'}`}></div>
                <span className="text-sm text-gray-600">
                  Availability: {overallAvailability === null ? 'No Data' : `${overallAvailability}%`}
                </span>
              </div>
            </div>
            <div className="flex items-center w-full relative mt-3">
              <div className="flex-1 flex items-center relative [&_*]:outline-none">
                <ResponsiveContainer width="100%" height={150}>
                  <BarChart
                    data={fullData}
                    barGap={0}
                    barCategoryGap={0}
                    margin={{ top: 0, right: 30, left: 30, bottom: 0 }}
                  >
                    <XAxis dataKey="date" hide />
                    <YAxis domain={[0, 100]} hide />
                    <Tooltip
                      cursor={{ fill: 'rgba(0,0,0,0.04)' }}
                      content={({ active, coordinate, payload }) => {
                        if (!active || !coordinate) return null;
                        
                        const dataPoint = payload?.[0]?.payload;
                        if (!dataPoint) return null;

                        const date = new Date(dataPoint.date);
                        const formattedDate = date.toLocaleDateString('zh-CN', {
                          year: '2-digit',
                          month: '2-digit',
                          day: '2-digit',
                        });

                        return (
                          <div className="rounded-lg shadow bg-white px-3 py-2 text-xs text-gray-800">
                            <div className="font-medium text-center border-b pb-1 mb-1">{formattedDate}</div>
                            {dataPoint.isEmpty ? (
                              <div className="text-gray-500 mt-1">No data</div>
                            ) : (
                              <>
                                <div className="mt-1">availability: <b>{dataPoint.availability}%</b></div>
                                <div>success/fail: <b>{dataPoint.succeed_count || '-'}/{dataPoint.failed_count || '-'}</b></div>
                                <div>p99: <b>{dataPoint.p99 !== undefined ? dataPoint.p99 + ' ms' : '-'}</b></div>
                              </>
                            )}
                          </div>
                        );
                      }}
                      isAnimationActive={false}
                    />
                    <Bar
                      dataKey={entry => entry.isEmpty ? 2 : entry.availability}
                      radius={[6, 6, 6, 6]}
                      minPointSize={2}
                      maxBarSize={16}
                      isAnimationActive={false}
                      name="availability"
                      background={{
                        fill: '#f3f4f6',
                        radius: 6
                      }}
                      className="group"
                      onMouseOver={(_, idx, e) => {
                        if (e && e.target) (e.target as HTMLElement).style.filter = 'brightness(1.1)';
                      }}
                      onMouseOut={(_, idx, e) => {
                        if (e && e.target) (e.target as HTMLElement).style.filter = '';
                      }}
                    >
                      {fullData.map((entry, index) => (
                        <Cell
                          key={`cell-${index}`}
                          fill={entry.isEmpty ? '#e2e4e9' : (planColors[plan] || '#4ade80')}
                          fillOpacity={entry.isEmpty ? 0.8 : 1}
                        />
                      ))}
                    </Bar>
                  </BarChart>
                </ResponsiveContainer>

              </div>
            </div>
            <div className="flex justify-between mt-2 text-sm text-gray-500" style={{ marginLeft: '30px', marginRight: '30px' }}>
              <span>{lastNDays[0]}</span>
              <span>Today</span>
            </div>
          </div>
        );
      })}
    </div>
  );
}

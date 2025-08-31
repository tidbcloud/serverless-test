"use client";
import React, { useEffect, useState } from 'react';
import RegionSelector from '@/components/RegionSelector';
import ProbeChart from '@/components/ProbeChart';

import { DateTime } from 'luxon';


function getLastNDatesWithUTC8(n: number) {
    const arr = [];
    const now = DateTime.now().setZone('UTC+8').startOf('day');
    for (let i = n - 1; i >= 0; i--) {
      const date = now.minus({ days: i });
      const formattedDate = date.toFormat('yyyy-MM-dd');
      arr.push(formattedDate);
    }
    return arr;
  }

export default function Home() {
  const [region, setRegion] = useState('');
  const [regions, setRegions] = useState<string[]>([]);
  const [allData, setAllData] = useState<any[]>([]);
  const [loading, setLoading] = useState(true);

  const filteredData = allData.filter(item => item.region === region);
  const lastNDays = getLastNDatesWithUTC8(60); 


  // 刷新数据的函数
  const refreshData = async () => {
    setLoading(true);
    try {
      // 并行获取region列表和所有数据
      const [regionsRes, dataRes] = await Promise.all([
        fetch('/api/probe-result/regions'),
        fetch('/api/probe-result')
      ]);
      
      const regionsList = await regionsRes.json();
      const allDataList = await dataRes.json();
      
      setRegions(regionsList);
      if (!region && regionsList.length > 0) {
        setRegion(regionsList[0]);
      }
      setAllData(allDataList);
    } catch (error) {
      console.error('Failed to fetch data:', error);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    refreshData();
  }, []);

  return (
    <main className="min-h-screen bg-gradient-to-b from-gray-50 to-white">
      <div className="container mx-auto px-4 py-12 max-w-7xl">
        <div className="space-y-8">
          <div style={{ display: 'flex', justifyContent: 'center', width: '100%' }}>
            <h1 className="text-4xl font-bold tracking-tight text-gray-900 sm:text-5xl">
              TiDB Cloud Connection Probe
            </h1>
          </div>
          
          {loading ? (
            <div className="fixed inset-0 bg-white/50">
              <div className="container mx-auto px-4 h-full flex items-center justify-center">
                <div style={{ display: 'flex', justifyContent: 'center', width: '100%' }}>
                  <span className="text-3xl font-bold text-gray-800">Loading data...</span>
                </div>
              </div>
            </div>
          ) : (
            <>
              <div style={{ display: 'flex', justifyContent: 'center', width: '100%' }}>
                <RegionSelector regions={regions} value={region} onChange={setRegion} />
              </div>
              <div className="mt-8">
                <ProbeChart data={filteredData} lastNDays={lastNDays}/>
              </div>
            </>
          )}
        </div>
      </div>
    </main>
  );
}

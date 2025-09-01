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


  const refreshData = async () => {
    console.info("refreshing data...");
    setLoading(true);
    try {
      const response = await fetch('/api/probe-result', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        cache: 'no-store'
      });
      const allDataList = await response.json();
      
      const uniqueRegions = Array.from(new Set(allDataList.map((item: any) => item.region as string))).sort() as string[];
      
      setRegions(uniqueRegions);
      if (!region && uniqueRegions.length > 0) {
        setRegion(uniqueRegions[0]);
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
          <div style={{ display: 'flex', justifyContent: 'right', width: '100%' }}>
              <button
                onClick={refreshData}
                className="px-4 py-2 bg-blue-500 text-white rounded hover:bg-blue-600 transition-colors"
                disabled={loading}
              >
                {loading ? 'Refreshing...' : 'Refresh Data'}
              </button>
          </div>
          
          {loading ? null : (
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

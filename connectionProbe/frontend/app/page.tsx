"use client";
import React, { useEffect, useState } from 'react';
import RegionSelector from '@/components/RegionSelector';
import ProbeChart from '@/components/ProbeChart';
import './globals.css';

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
        <div className="space-y-10">
          <div className="flex justify-center w-full">
            <h1 className="relative text-4xl font-normal tracking-tight text-gray-700 sm:text-5xl pb-3">
              TiDB Cloud Probe
              <span className="absolute bottom-0 left-0 w-full h-px bg-gray-200"></span>
            </h1>
          </div>
          {loading ? (
            <div className="flex justify-center items-center min-h-[calc(100vh-250px)]">
              <div className="animate-spin rounded-full h-12 w-12 border-4 border-dashed border-gray-300 border-t-transparent"></div>
            </div>
          ) : (
            <>
              <div className="flex items-center justify-between mb-8">
                <RegionSelector regions={regions} value={region} onChange={setRegion} />
                <button
                  onClick={refreshData}
                  className="px-3 py-1 bg-blue-500 text-white text-sm rounded hover:bg-blue-600 transition-colors"
                  disabled={loading}
                >
                  Refresh
                </button>
              </div>
              <div>
                <ProbeChart data={filteredData} lastNDays={lastNDays}/>
              </div>
            </>
          )}
        </div>
      </div>
    </main>
  );
}

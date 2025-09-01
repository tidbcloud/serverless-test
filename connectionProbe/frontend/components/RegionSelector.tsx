import * as React from 'react';
import * as Select from '@radix-ui/react-select';
import { ChevronDownIcon } from 'lucide-react';

interface Props {
  regions: string[];
  value: string;
  onChange: (value: string) => void;
}

export default function RegionSelector({ regions, value, onChange }: Props) {
  return (
    <div className="flex justify-end w-full mb-8">
      <div className="flex items-center gap-3">
        <span className="font-medium text-gray-700 text-sm">region </span>
        <Select.Root value={value} onValueChange={onChange}>
          <Select.Trigger
            className="inline-flex h-8 items-center gap-1 rounded-md border border-gray-200 bg-white px-3 text-sm font-medium text-gray-800 shadow-sm transition-all hover:bg-gray-50/70 focus:outline-none focus:ring-2 focus:ring-blue-100 min-w-[140px]"
            aria-label="Region"
          >
            <div className="flex items-center justify-between w-full">
              <Select.Value />
              <Select.Icon>
                <ChevronDownIcon style={{ height: '1rem', width: '1rem', color: '#6B7280', transition: 'transform 0.2s' }} 
                  className="data-[state=open]:rotate-180" />
              </Select.Icon>
            </div>
          </Select.Trigger>
          <Select.Portal>
            <Select.Content
              className="z-50 min-w-[140px] overflow-hidden rounded-lg border border-gray-100 bg-white text-gray-800 shadow-lg"
              position="popper"
              sideOffset={4}
            >
              <div className="max-h-[180px] overflow-y-auto py-1 px-1">
                {regions.map((region) => (
                  <Select.Item
                    key={region}
                    value={region}
                    style={{ backgroundColor: '#F3F4F6', color: '#111827' }}
                    className="flex w-full select-none items-center rounded-md py-1.5 px-2 text-sm font-medium outline-none cursor-pointer"
                  >
                    <Select.ItemText>{region}</Select.ItemText>
                  </Select.Item>
                ))}
              </div>
            </Select.Content>
          </Select.Portal>
        </Select.Root>
      </div>
    </div>
  );
}

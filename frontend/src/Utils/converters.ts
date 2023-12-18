const TBinMB = 1024 * 1024;
const GBinMB = 1024;

export const formatMemory = (mb: number): string => {
  if (mb < GBinMB) {
    return `${mb} MB`;
  }
  if (mb < TBinMB) {
    return `${Math.round(mb / GBinMB)} GB`;
  }
  return `${Math.round(mb / TBinMB)} TB`;
};

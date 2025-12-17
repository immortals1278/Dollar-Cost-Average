import mainnet from "../data/addresses.mainnet.json";
import testnet from "../data/addresses.testnet.json";

// Ensure that mainnet and testnet are typed as arrays
const mainnetData: any[] = mainnet as any[];
const testnetData: any[] = testnet as any[];

const extractUniqueValues = (data: any[], key: string): string[] => {
  const allValues = data.filter((item) => item[key] !== undefined).map((item) => item[key].toString());
  return [...new Set(allValues)];
};

const generateTypesForKeys = (keys: string[]) => {
  const networks = [...mainnetData, ...testnetData];

  let typeDefs = "";

  keys.forEach((key) => {
    const uniqueValues = extractUniqueValues(networks, key);
    if (uniqueValues.length > 0) {
      const isNumeric = uniqueValues.every((value) => !isNaN(Number(value)));
      const formattedValues = isNumeric ? uniqueValues.join(" | ") : `"${uniqueValues.join('" | "')}"`;

      typeDefs += `export type Param${toCamelCase(key, true)} = ${formattedValues};\n`;
    }
  });

  console.log(typeDefs);
};

// Modify this function to handle underscores and capitalize each word
const toCamelCase = (string: string, capitalizeFirst: boolean = false) => {
  return string
    .split("_")
    .map((word, index) => {
      if (index === 0) {
        return capitalizeFirst ? word.charAt(0).toUpperCase() + word.slice(1) : word;
      }
      return word.charAt(0).toUpperCase() + word.slice(1);
    })
    .join("");
};

const keysToGenerate = ["symbol", "chain_name", "type"];

generateTypesForKeys(keysToGenerate);

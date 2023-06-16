class Car {
  constructor(brand, color, maxSpeed, chassisNumber) {
    this.brand = brand;
    this.color = color;
    this.maxSpeed = maxSpeed;
    this.chassisNumber = chassisNumber;
  }
 
  drive() {
    console.log(`${this.brand} ${this.color} is driving`);
  }
 
  reverse() {
    console.log(`${this.brand} ${this.color} is reversing`);
  }
 
  turn() {
    console.log(`${this.brand} ${this.color} is turning`);
  }
}
 
// Membuat objek mobil dengan constructor function Car
const car1 = new Car('Toyota', 'Silver', 200, 'to-1');
const car2 = new Car('Honda', 'Black', 180, 'ho-1');
const car3 = new Car('Suzuki', 'Red', 220, 'su-1');
 
// console.log(car1);
// console.log(car2);
// console.log(car3);
 
// car1.drive();
// car2.turn();
// car3.reverse();

class Spesifikasi {

  constructor(prosesor, ram, motherboard, ssd) {
    this.prosesor = prosesor;
    this.ram = ram;
    this.motherboard = motherboard;
    this.ssd = ssd;
  }
  pamer(){
    console.log(`pc saya memakai ${this.prosesor} sebagai prosesornya dan ram merk ${this.ram} 2*8gb`);
  }
  pamerLagi(){
    console.log(`kalo mobonya pake ${this.motherboard} aja si sama ssdnya merk ${this.ssd}`);
  }
}

class SpesifikasiIntel extends Spesifikasi{
  pamer(){
    console.log(`ASDAS DAS AS ${this.prosesor} ASD SAD ASD SAD SACSACSA C S ${this.ram} 2*8gb`);
  }
  pamerLagi(){
    console.log(`ASD AS ASDSAD SAD SAD ${this.motherboard} SAC ASC SACWCQWD W ${this.ssd}`);
  }
}

const amd = new Spesifikasi("amd ryzen 5", 'fury', 'asrock', 'teamgroup');
console.log(amd);
amd.pamer();
amd.pamerLagi();

let intel = new SpesifikasiIntel('intel i5', 'kingston' , 'asus b450', 'samsung');

intel.pamer();
intel.pamerLagi();





// class Person {
//   constructor(name) {
//     console.info(`Membuat person ${name}`);
//   }
// }

// class Animal extends Person{

// }

// const syahrul = new Person('Syahrul');
// console.info(syahrul);

// const beruang = new Animal('beruang');
// console.info(beruang);
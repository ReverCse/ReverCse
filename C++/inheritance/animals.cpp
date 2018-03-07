#include <iostream>
#include <vector>

class Animal {
public:
  Animal() : name("none"), weight(0) {}
  Animal(std::string str, int w) : name(str), weight(w) {}
  virtual std::string getSound() = 0; // what sound does the animal make?
  friend std::ostream& operator<<(std::ostream& os, const Animal& a) {
    return os << a.name << " weighs " << a.weight <<"lbs";
  }
protected:
  std::string name;
  int weight;
};

class Fox : public Animal {
public:
  Fox() : Animal() {}
  // why reinvent the wheel? We already wrote this code. Just call it again
  Fox(std::string str, int w) : Animal(str, w) {}
  virtual std::string getSound() { return "I'M A FOX WOO!"; }

  // weight is protected. Derived classes can touch and manipulate the weight
  int loseWeight() {return --weight;}
};

class OrangeFox : public Fox {
public:
  OrangeFox() : Fox() {} // We're going deeper
  OrangeFox(std::string str, int w) : Fox(str, w) {}
  std::string getSound() { return "I'm orange!"; }
};
int main() {
  // Animal animal; <-- ERROR cannot instantiate a class with pure virtual function
  Fox fox("Rupert", 22);
  std::cout << fox << std::endl;
  std::cout << "RUPERT! GO ON A DIET!\n";
  fox.loseWeight();
  std::cout << fox << std::endl;
  std::cout << fox.getSound() << std::endl << std::endl;

  
  OrangeFox oFox("Stanley", 10);
  std::cout << oFox << std::endl;
  std::cout << "STANLEY! GET SMALL!\n";
  oFox.loseWeight();
  std::cout << oFox << std::endl;
  std::cout << oFox.getSound() << std::endl;
  std::cout << static_cast<Fox>(oFox).getSound() << std::endl;
}

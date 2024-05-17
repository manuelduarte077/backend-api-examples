defmodule ValidateTest do
    use ExUnit.Case
    doctest Shop.Validate

    describe "get_required/2" do
      test "returns {:ok, value} when field is present" do
        params = %{name: "John"}
        assert {:ok, "John"} = Shop.Validate.get_required(params, :name)
      end

      test "returns {:error, message} when field is nil" do
        params = %{name: nil}
        assert {:error, "Field name is required"} = Shop.Validate.get_required(params, :name)
      end
    end

    describe "is_integer/2" do
      test "returns {:ok, integer} when value is an integer" do
        assert {:ok, 42} = Shop.Validate.is_integer("42", :age)
      end

      test "returns {:error, message} when value is not an integer" do
        assert {:error, "age: Value is not an integer"} = Shop.Validate.is_integer("abc", :age)
      end
    end
  end
